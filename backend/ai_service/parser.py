import json
import re

def parse_openai_response_to_json(response: str) -> dict:
    """
    Parse a JSON-like string response from OpenAI into a valid dictionary.
    Handles cases with code block delimiters and poorly formatted JSON.
    """
    
    if response.startswith("```") and response.endswith("```"):
        response = response.strip("```").strip()
        
    result = {}
    try:
        for line in response.splitlines():
            line = line.strip()
            if not line or ":" not in line:
                continue
            key, value = map(str.strip, line.split(":", 1))
            
            value = value.strip('"').strip("'")
            result[key] = value
    except Exception as e:
        raise ValueError(f"Failed to parse response: {e}")

    return result