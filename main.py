from fastapi import FastAPI, HTTPException
from enum import Enum
from pydantic import BaseModel

app = FastAPI()

class DataTypes(str, Enum):
    CSV = 'csv'

class DataItemBase(BaseModel):
    type: DataTypes
    data: bytes

class Connector:
    def __init__(self, type: DataTypes):
        self.type = type

    def connect(self):
        if self.type is DataTypes.CSV:
            return self.__csv_connector()
        else:
            # TODO: create custom Exceptions
            raise Exception("invalid connector")




@app.get("/")
async def root():
    return {"message": "Hello from Leyndash!"}

@app.post("/api/{data}")
async def addDataConnector(data: DataItemBase):
    if data.type is DataTypes.CSV:
        #? How do we model ingested data
        pass
    else:
        # Need to implement other data connectors
        # Throw error for now
        raise HTTPException(404, details= "Undefined connector")

