from fastapi import FastAPI
from enum import Enum

app = FastAPI()

class DataTypes(str, Enum):
    CSV = 'csv'

class DataItem:
    type: DataTypes
    data: bytes

@app.get("/")
async def root():
    return {"message": "Hello from Leyndash!"}

@app.post("/api/{data}")
async def addDataConnector(data: DataItem):
    if data.type is DataTypes.CSV:
        #? How do we model ingested data
        pass
    else:
        # Need to implement other data connectors
        # Throw error for now
        return {"failure": "Undefined connector"}

