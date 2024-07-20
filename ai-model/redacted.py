from flask import Flask, request, jsonify, render_template
import joblib
import pandas as pd
from datetime import datetime

app = Flask(__name__)

# Load the models and encoder
model_demand = joblib.load("model_demand.pkl")
model_stock = joblib.load("model_stock.pkl")
encoder = joblib.load("encoder.pkl")

# Categorical features used in encoding
categorical_features = ["Blood Type", "Hospital"]


# Function to predict blood demand and stock level
def predict_blood_demand_and_stock(date, blood_type, hospital):
    date = pd.to_datetime(date)
    year, month, day = date.year, date.month, date.day

    input_data = {
        "Year": year,
        "Month": month,
        "Day": day,
    }

    for col in encoder.get_feature_names_out(categorical_features):
        input_data[col] = 0

    input_data[f"Blood Type_{blood_type}"] = 1
    input_data[f"Hospital_{hospital}"] = 1

    input_df = pd.DataFrame([input_data])
    predicted_demand = model_demand.predict(input_df)[0]
    predicted_stock = model_stock.predict(input_df)[0]

    return predicted_demand, predicted_stock


@app.route("/predictBloodDemand")
def index():
    return render_template("demand.html")


@app.route("/bloodRequest", methods=["GET", "POST"])
def bloodRequest():
    if request.method == "POST":
        recipient = {
            "blood_type": request.form["blood_type"],
            "location": request.form["location"],
            "urgency": int(
                request.form["urgency"]
            ),  # Assume this is directly usable in scoring if needed
        }

        donors = pd.read_csv("updated_donor_dataset.csv")
        top_donors = calculate_score(donors, recipient)
        return render_template(
            "result2.html", donors=top_donors.to_dict(orient="records")
        )
    return render_template("request.html")

@app.route("/predict", methods=["POST"])
def predict():
    date = request.form["date"]
    blood_type = request.form["blood_type"]
    hospital = request.form["hospital"]

    predicted_demand, predicted_stock = predict_blood_demand_and_stock(
        date, blood_type, hospital
    )

    return render_template(
        "demand.html", predicted_demand=predicted_demand, predicted_stock=predicted_stock
    )


def calculate_score(donors, recipient):
    # Example mapping from location names to numeric codes
    location_weights = {
        "Port Louis": 1,
        "Curepipe": 2,
        "Flic en Flac": 3,
        "Vacoas": 4,
        "Goodlands": 5,
        "Rose Hill": 6,
    }

    # Blood type compatibility mapping
    compatible_blood = {
        "O-": ["O-"],
        "O+": ["O+", "O-"],
        "A-": ["A-", "O-"],
        "A+": ["A+", "A-", "O+", "O-"],
        "B-": ["B-", "O-"],
        "B+": ["B+", "B-", "O+", "O-"],
        "AB-": ["AB-", "A-", "B-", "O-"],
        "AB+": ["AB+", "AB-", "A+", "A-", "B+", "B-", "O+", "O-"],
    }

    # Ensure recipient location is converted to integer
    recipient_location_code = location_weights.get(recipient["location"], None)
    if recipient_location_code is None:
        print("Location code not found for:", recipient["location"])
        return pd.DataFrame()  # Return an empty DataFrame if location code is not found

    donors["Location Code"] = donors["Location"].map(location_weights)
    donors["Location Code"] = pd.to_numeric(
        donors["Location Code"], errors="coerce"
    )  # Ensure numeric

    # Filter donors by compatible blood types
    if recipient["blood_type"] not in compatible_blood:
        print("Invalid blood type for recipient:", recipient["blood_type"])
        return pd.DataFrame()
    filtered_donors = donors[
        donors["Blood Type"].isin(compatible_blood[recipient["blood_type"]])
    ]

    # Weighted criteria
    weights = {
        "blood_type_compatibility": 0.5,
        "location_proximity": 0.3,
        "donation_recency": 0.2,
    }

    # Blood type compatibility score with higher weight for exact matches
    def blood_type_score(blood_type, recipient_type):
        if blood_type == recipient_type:
            return 10  # Exact match
        elif blood_type in compatible_blood[recipient_type]:
            return 8  # Compatible but not exact
        else:
            return 0  # Not compatible

    filtered_donors["blood_type_score"] = filtered_donors["Blood Type"].apply(
        lambda x: blood_type_score(x, recipient["blood_type"])
    )

    # Location proximity score
    filtered_donors["location_score"] = 10 - abs(
        filtered_donors["Location Code"] - recipient_location_code
    )

    # Convert 'Donation History' to datetime and calculate days since last donation
    filtered_donors["Donation History"] = pd.to_datetime(
        filtered_donors["Donation History"], errors="coerce"
    )
    filtered_donors["days_since_last_donation"] = (
        pd.Timestamp("now") - filtered_donors["Donation History"]
    ).dt.days
    filtered_donors["recency_score"] = 100 / (
        filtered_donors["days_since_last_donation"] + 1
    )  # More recent donations get higher score

    # Calculate total score with weights
    filtered_donors["Total Score"] = (
        filtered_donors["blood_type_score"] * weights["blood_type_compatibility"]
        + filtered_donors["location_score"] * weights["location_proximity"]
        + filtered_donors["recency_score"] * weights["donation_recency"]
    )

    return filtered_donors.nlargest(5, "Total Score")


if __name__ == "__main__":
    app.run(debug=True)
