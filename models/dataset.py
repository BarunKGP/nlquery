"""Prepare a ChunkDataset to enable streaming chunks of data.
Reference: https://huggingface.co/learn/cookbook/en/fine_tuning_code_llm_on_single_gpu
"""

import torch
from tqdm import tqdm

from transformers import (
    AutoModelForCausalLM,
    AutoTokenizer,
    Trainer,
    TrainingArguments,
    logging,
    set_seed,
    BitsAndBytesConfig,
)

import SEED from config

set_seed(SEED)

#> Load Dataset
# dataset = load_dataset(
#     DATASET,
#     data_dir="data",
#     split="train",
#     streaming=True,
# )

valid_data = dataset.take(4000)
train_data = dataset.skip(4000)
train_data = train_data.shuffle(buffer_size=5000, seed=SEED)

tokenizer = AutoTokenizer.from_pretrained(model, trust_remote_code=True)