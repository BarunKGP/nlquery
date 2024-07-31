import os

from dotenv import load_dotenv

from main import Connector, DataItemBase
import torch
import pandas as pd

from transformers import pipeline

MODEL_ID = "meta-llama/Meta-Llama-3.1-8B-Instruct"


class PlotlyModel:
    def __init__(self, model_id, device="cpu"):
        self.model_id = model_id

        load_dotenv()

        with open("data/_example_claude_plotly.html", "r") as f:
            example_plotly_code = f.read()

        plotly_prompt = """You are a helpful coding assistant that helps users visualize their data using Plotly.js.
        You will be provided the user's visualization query and the schema of the CSV file as a JSON object.
        You should then return the relevant visualization code in Plotly.
        Take you time to think through the problem step by step and return a Plotly.js code snippet that visualizes the data
        You must return the code snippet as a JSON object with the key `plotlyCode`.
        You must return only JSON and no other text. "
        
        Here is an example:
        Input: 
        {
            "user": "Generate a bar chart displaying the total revenue over the last three months",
            "schema": {
                "filename": "revenue_last3m.csv",
                "columns": {
                    "id": "string",
                    "item_id": "string",
                    "date": "datetime",
                    "units_sold": "int32",
                    "price_per_unit": "float64"
                }
            }
        }

        Response: {
            "plotlyCode":"""

        self.system_prompt = plotly_prompt + f" {example_plotly_code} \n" + "}"
        self.device = device

    def run(self, user_content, connector):
        file_schema = connector.infer_schema()
        user_prompt = {"user": user_content.prompt, "schema": file_schema}
        messages = [
            {"role": "system", "content": self.system_prompt},
            {"role": "user", "content": f"{user_prompt}"},
        ]

        pipe = pipeline(
            "text-generation",
            model=self.model_id,
            model_kwargs={
                "torch_dtype": torch.bfloat16,
            },  # "quantization_config": {"use_in_4bit": True}},
            device=self.device,
        )

        print("Getting outputs...")
        outputs = pipe(
            messages,
            max_new_tokens=512,
        )

        print("Received outputs. Printing response...")
        assistant_response = outputs[0]["generated_text"][-1]["content"]
        # print(assistant_response)

        return assistant_response
