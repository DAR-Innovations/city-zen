import openai
from typing import List

openai.api_key = "YOUR_OPENAI_API_KEY"

def generate_task_metadata(image_url: str, issue_description: str, departments: dict) -> dict:
    department_list = "\n".join([f"{dept}: {', '.join(tasks)}" for dept, tasks in departments.items()])

    prompt = f"""
        Analyze the following issue based on the image and description:

        Issue Description: {issue_description}
        Image URL: {image_url}

        Departments and Task Types:
        {department_list}

        Determine:
        1. If the task should go to Volunteers or a specific Department.
        2. The urgency of the task (LOW, MEDIUM, HIGH).
        3. The complexity of the task (LOW, MEDIUM, HIGH).
        4. Generate a suitable title and description for the task.

        Respond in JSON format:
        {{
            "title": "...",
            "description": "...",
            "urgency": "...",
            "complexity": "...",
            "department": "..." // null if for volunteers
        }}
        """
    response = openai.completions.create(
        model="gpt-4",
        messages=[{"role": "user", "content": prompt}]
    )
    return response["choices"][0]["message"]["content"]


def validate_task_completion(original_image_url: str, completed_image_url: str) -> dict:
    prompt = f"""
        Compare the following two images:

        1. Original Image: {original_image_url}
        2. Completed Image: {completed_image_url}

        Determine if the task shown in the original image is resolved in the completed image.
        Respond in JSON format:
        {{
            "is_valid": true/false,
            "reason": "..."
        }}
        """
    response = openai.ChatCompletion.create(
        model="gpt-4",
        messages=[{"role": "user", "content": prompt}]
    )
    return response["choices"][0]["message"]["content"]
