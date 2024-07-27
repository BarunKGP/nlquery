import os

import torch
import pandas as pd

from transformers import pipeline

MODEL_ID = "meta-llama/Meta-Llama-3.1-8B-Instruct"


def init_model():
    os.environ["HF_TOKEN"] = "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"
    os.environ["HUGGINGFACEHUB_API_TOKEN"] = "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"

    with open('data/_example_claude_plotly.html', 'r') as f:
        example_plotly_code = f.read()

    PLOTLY_PROMPT = """
    You are a helpful coding assistant that helps users visualize their data using Plotly.js.
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

    PLOTLY_PROMPT += f' {example_plotly_code} \n' + '}'
    return PLOTLY_PROMPT


def construct_user_prompt(csv_file):
    df = pd.read_csv(csv_file)
    
    col_schema = {}
    for c in df.columns:
        col_type = str(df[c].dtype)
        if col_type == "object":
            col_type = 'datetime' if 'date' in c else 'string'

        col_schema[c] = col_type


    # col_schema = {c: str(df[c].dtype) for c in df.columns}
    return {"filename": os.path.basename(csv_file), "columns": col_schema}


def main():
    # user_prompt = input("Enter your nlQuery:")
    # csv_file = input("Enter CSV file path:")
    user_prompt = "Generate a bar chart showing the average temperature each week"
    csv_file = "data/revenue_last3m.csv"

    PLOTLY_PROMPT = init_model()
    user_content = {"user": user_prompt, "schema": construct_user_prompt(csv_file)}

    # print(f'Plotly prompt: {PLOTLY_PROMPT}')
    # print(f'\nUser prompt: {user_content}')
    messages = [
        {
            "role": "system", 
            "content": PLOTLY_PROMPT
        },
        {
            "role": "user", 
            "content": f"{user_content}"
        },
    ]

    print('Initialized prompts')

    pipe = pipeline(
        "text-generation",
        model=MODEL_ID,
        model_kwargs={"torch_dtype": torch.bfloat16,}, # "quantization_config": {"use_in_4bit": True}},
        device="cuda",
    )

    print("Getting outputs...")
    outputs = pipe(
        messages,
        max_new_tokens=512,
    )

    print('Received outputs. Printing response...')
    assistant_response = outputs[0]["generated_text"][-1]["content"]
    print(assistant_response)

if __name__ == '__main__':
    main()
