import spacy
from spacy.language import Language
from spacy.matcher import Matcher
from spacy.tokens import Doc

# Load the spaCy model
nlp = spacy.load("en_core_web_sm")

# Register custom attributes
Doc.set_extension("chart_type", default=None, force=True)
Doc.set_extension("time_frame", default=None, force=True)

# Define custom components
@Language.component("add_chart_type")
def add_chart_type(doc):
    matcher = Matcher(nlp.vocab)
    patterns = [
        [{"LOWER": {"IN": ["bar", "line", "pie", "scatter", "histogram"]}}, 
         {"LOWER": {"IN": ["chart", "graph", "plot"]}, "OP": "?"}],
        [{"LOWER": "horizontal"}, {"LOWER": "bar"}, 
         {"LOWER": {"IN": ["chart", "graph", "plot"]}, "OP": "?"}]
    ]
    matcher.add("CHART_TYPE", patterns)
    matches = matcher(doc)
    if matches:
        start, end = matches[0][1], matches[0][2]
        doc._.chart_type = doc[start:end].text
    return doc

@Language.component("add_time_frame")
def add_time_frame(doc):
    matcher = Matcher(nlp.vocab)
    patterns = [
        [{"LOWER": {"IN": ["last", "past", "previous"]}}, 
         {"LIKE_NUM": True}, 
         {"LOWER": {"IN": ["day", "week", "month", "year"]}, "OP": "+"}]
    ]
    matcher.add("TIME_FRAME", patterns)
    matches = matcher(doc)
    if matches:
        start, end = matches[0][1], matches[0][2]
        doc._.time_frame = doc[start:end].text
    return doc

# Add custom components to the pipeline
nlp.add_pipe("add_chart_type", last=True)
nlp.add_pipe("add_time_frame", last=True)

def analyze_query(query):
    doc = nlp(query)
    
    result = {
        "chart_type": doc._.chart_type,
        "time_frame": doc._.time_frame,
        "colors": [],
        "styles": [],
        "measures": [],
        "dimensions": []
    }
    
    color_terms = ["blue", "red", "green", "yellow", "purple", "orange", "pink", "brown", "gray", "black", "white", "navy", "maroon"]
    style_terms = ["bold", "italic", "underline", "strikethrough", "large", "small", "thick", "thin"]
    
    for token in doc:
        if token.pos_ == "NOUN" and token.dep_ in ["dobj", "pobj", "attr"]:
            if token.text in ["revenue", "sales", "profit"]:  # Extend this list as needed
                result["measures"].append(token.text)
            elif token.text not in ["chart", "graph", "plot"]:
                result["dimensions"].append(token.text)
        elif token.pos_ == "ADJ" or token.pos_ == "PROPN":  # Include proper nouns for color names
            if token.text.lower() in color_terms:
                result["colors"].append(token.text.lower())
            elif token.text.lower() in style_terms:
                result["styles"].append(token.text.lower())
    
    # Clean up dimensions
    result["dimensions"] = [dim for dim in result["dimensions"] if dim not in ["bar", "bars", "tick", "ticks"]]
    
    return result

def main():
    query = input("Enter query: ")
    result = analyze_query(query)
    print(result)

if __name__ == "__main__":
    main()
