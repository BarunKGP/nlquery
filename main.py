from fastapi import FastAPI, HTTPException
from enum import Enum
from pydantic import BaseModel

app = FastAPI()

class DataTypes(str, Enum):
    CSV = 'csv'

class DataItemBase(BaseModel):
    type: DataTypes
    data: bytes

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

