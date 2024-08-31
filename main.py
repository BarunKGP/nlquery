from enum import Enum
import json
import os

from fastapi import FastAPI, HTTPException
import pandas as pd
from pydantic import BaseModel

from models.plotly_model import PlotlyModel

app = FastAPI()


class DataTypes(str, Enum):
    CSV = "csv"


class DataItemBase(BaseModel):
    data_type: DataTypes
    data: bytes
    prompt: str


class newClass:
    def __init__(self, name):
        self.name = name
        a = [
            "Several",
            "words",
            "together",
            "in",
            "a",
            "really,",
            "annoyingly",
            "long",
            "sentence.", "I'm", "stumped,", "istg!",
        ]
        print(
            "This is a really long string"
            + "Will the LSP correctly split this?"
            + "Come on! Split this already!"
            + "Will you??!" + "I am testing out if the LSP formatting is correct" + "testing even more!" + "another test"
        )

        print(" ".join(a))


class Connector:

    def __init__(self, type: DataTypes, data_file):
        self.type = type
        self.data = self.connect(data_file)
        self.filename = os.path.basename(data_file)

    def _csv_connector(self, data):
        df = pd.DataFrame.from_csv(data)
        return df

    def connect(self, data):
        """Generates a connector for `data` so that we can interact with it and
        process it

        Args:
            data (_type_): data to connect to

        Raises:
            Exception: _description_

        Returns:
            _type_: _description_
        """

        if self.type is DataTypes.CSV:
            return self._csv_connector(data)
        # TODO: create custom Exceptions
        raise HTTPException(404, "invalid connector type")

    def infer_schema(self):
        col_schema = {}
        for c in self.data.columns:
            col_type = str(self.data[c].dtype)
            if col_type == "object":
                col_type = "datetime64" if "date" in c else "string"

            col_schema[c] = col_type

        # col_schema = {c: str(self.data[c].dtype) for c in self.data.columns}
        return {"filename": self.filename, "columns": col_schema}


class PlotlyResponse(BaseModel):
    plotlyCode: str


@app.get("/")
async def root():
    return {"message": "Hello from Leyndash!"}


@app.post("/api/v1/prompt/{data}")
async def addDataConnector(data: DataItemBase) -> PlotlyResponse:
    if not data or data.type.lower() != "csv":
        # Need to implement other data connectors
        # Throw error for now
        raise HTTPException(404, detail="Undefined connector")

    # Initialize data connector
    connector = Connector(type=data.type, data_file=data.contents)

    # Init model and pipeline
    try:
        model = PlotlyModel("meta-llama/Meta-Llama-3.1-8B-Instruct", device="cuda")
    except Exception as e:
        raise HTTPException(500, detail="Error initializing model") from e

    # Pass data to model
    try:
        model_response = model.run(data, connector)
    except Exception as e:
        raise HTTPException(500, detail="Error obtaining model response") from e

    # ? Validate that model response?
    # Parse string to JSON
    try:
        model_response = json.loads(model_response)
    except Exception as e:
        raise HTTPException(406, detail="Error parsing model response to JSON") from e

    # Pass outputs back to ui
    return model_response
