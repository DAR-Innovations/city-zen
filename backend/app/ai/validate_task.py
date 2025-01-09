import openai
from fastapi import HTTPException, UploadFile
from app.core.config import settings

openai.api_key = settings.OPENAI_API_KEY

async def validate_task_completion(file: UploadFile, task):
    """
    Validate task completion using OpenAI API.
    :param file: Uploaded file of the completed task (image).
    :param task: Task object containing issue details.
    :return: True if validation passes, False otherwise.
    """
    try:
        # Read the file content
        image_content = await file.read()

        # Use OpenAI API to validate the completion
        prompt = f"""
        You are an AI trained to validate task completion based on visual evidence and descriptions.
        Task Description: {task.description}
        Task Urgency: {task.urgency}
        Task Complexity: {task.complexity}

        Does the attached image indicate the task is completed successfully? Provide a simple "yes" or "no".
        """

        # Use the OpenAI API (mock the file upload for now, as OpenAI API does not directly process images)
        response = openai.ChatCompletion.create(
            model="gpt-4",
            messages=[
                {"role": "system", "content": "You are an AI validator."},
                {"role": "user", "content": prompt}
            ],
        )

        result = response["choices"][0]["message"]["content"].strip().lower()
        return result == "yes"
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Failed to validate task completion: {str(e)}")
