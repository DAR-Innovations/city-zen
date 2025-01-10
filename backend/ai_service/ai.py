import os
import openai
from openai import OpenAI
from dotenv import load_dotenv
from parser import parse_openai_response_to_json

load_dotenv()

API_KEY = os.getenv("OPENAI_API_KEY")

openai.api_key = API_KEY
client = OpenAI(api_key=API_KEY)

def generate_task_metadata(image_url: str, issue_description: str, departments: dict) -> dict:
    """
    Generate task metadata using OpenAI's structured messages.
    """
    department_list = [{"type": "text", "text": f"{dept}: {', '.join(tasks)}"} for dept, tasks in departments.items()]

    response = client.chat.completions.create(
        model="gpt-4o-mini",
        messages=[
            {
                "role": "user",
                "content": [
                    {"type": "text", "text": "Analyze the following issue based on the image and description:"},
                    {"type": "text", "text": f"Issue Description: {issue_description}"},
                    {"type": "image_url", "image_url": {"url": image_url}},
                    {"type": "text", "text": "Departments and Task Types:"},
                    *department_list,
                    {"type": "text", "text": """
                        Determine:
                        1. If the task should go to Volunteers or a specific Department.
                        2. The urgency of the task (LOW, MEDIUM, HIGH).
                        3. The complexity of the task (LOW, MEDIUM, HIGH).
                        4. Generate a suitable title and description for the task.

                        Respond in JSON format:
                        {
                            "title": "...",
                            "description": "...",
                            "urgency": "...",
                            "complexity": "...",
                            "department": "..." // null if for volunteers
                        }
                    """},
                ],
            }
        ],
        max_tokens=300,
    )

    raw_content = response.choices[0].message.content
    return parse_openai_response_to_json(raw_content)



def validate_task_completion(original_image_url: str, completed_image_url: str) -> dict:
    """
    Validate task completion by comparing two images using OpenAI's structured messages.
    """
    response = client.chat.completions.create(
        model="gpt-4o-mini",
        messages=[
            {
                "role": "user",
                "content": [
                    {"type": "text", "text": "Compare the following two images:"},
                    {"type": "text", "text": "1. Original Image:"},
                    {"type": "image_url", "image_url": {"url": original_image_url}},
                    {"type": "text", "text": "2. Completed Image:"},
                    {"type": "image_url", "image_url": {"url": completed_image_url}},
                    {"type": "text", "text": """
                        Determine if the task shown in the original image is resolved in the completed image.
                        Respond in JSON format:
                        {
                            "is_valid": true/false,
                            "reason": "..."
                        }
                    """},
                ],
            }
        ],
        max_tokens=300,
    )

    return response.choices[0].message.content
