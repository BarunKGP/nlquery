from enum import Enum

from fastapi import FastAPI, HTTPException
import pandas as pd
from pydantic import BaseModel

app = FastAPI()


class DataTypes(str, Enum):
    CSV = "csv"


class DataItemBase(BaseModel):
    type: DataTypes
    data: bytes


class Connector:
    def __init__(self, type: DataTypes):
        self.type = type

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
        raise Exception("invalid connector")


@app.get("/")
async def root():
    return {"message": "Hello from Leyndash!"}


@app.post("/api/{data}")
async def addDataConnector(data: DataItemBase):
    if data.type is DataTypes.CSV:
        # ? How do we model ingested data
        pass
    else:
        # Need to implement other data connectors
        # Throw error for now
        raise HTTPException(404, detail="Undefined connector")
