from sqlalchemy.orm import Session
from app.models.tasks.task_types import TaskTypes
from app.models.departments.departments import Departments
from app.models.departments.department_task_types import DepartmentTaskTypes
import openai
from app.core.config import settings

openai.api_key = settings.OPENAI_API_KEY

def generate_task_metadata(image_path: str, metadata_path: str, db: Session) -> dict:
    """
    Generate task metadata using image and metadata analysis, including title and description.
    :param image_path: Path to the uploaded image.
    :param metadata_path: Path to the associated metadata file.
    :param db: Database session.
    :return: Dictionary containing task metadata (title, description, urgency, complexity, department).
    """
    try:
        # Read metadata
        with open(metadata_path, "r") as f:
            metadata = f.read()

        # Retrieve department and task type data
        departments = db.query(Departments).all()
        department_info = []
        for department in departments:
            task_types = (
                db.query(TaskTypes.name)
                .join(DepartmentTaskTypes, DepartmentTaskTypes.task_type_id == TaskTypes.id)
                .filter(DepartmentTaskTypes.department_id == department.id)
                .all()
            )
            department_info.append(
                {
                    "department_name": department.name,
                    "task_types": [t.name for t in task_types],
                }
            )

        # Use OpenAI API to analyze metadata and generate task details
        prompt = f"""
        Analyze the following image metadata and determine the task details:
        Metadata: {metadata}

        Departments and Their Capabilities:
        {department_info}

        Provide:
        1. A suitable title for the task.
        2. A detailed description of what needs to be done.
        3. The urgency (Low, Medium, High).
        4. The complexity (Low, Medium, High).
        5. If it's High complexity, assign it to the most appropriate department.
        """

        response = openai.ChatCompletion.create(
            model="gpt-4",
            messages=[
                {"role": "system", "content": "You are an AI trained to analyze metadata."},
                {"role": "user", "content": prompt},
            ],
        )

        result = response["choices"][0]["message"]["content"]

        # Extract relevant fields from the response
        title = result.split("Title:")[1].split("\n")[0].strip()
        description = result.split("Description:")[1].split("\n")[0].strip()
        urgency = "High" if "High" in result else "Medium" if "Medium" in result else "Low"
        complexity = "High" if "High" in result else "Medium" if "Medium" in result else "Low"
        department = (
            result.split("Department:")[1].strip() if "Department:" in result else None
        )

        return {
            "title": title,
            "description": description,
            "urgency": urgency,
            "complexity": complexity,
            "department": department,
        }

    except Exception as e:
        raise ValueError(f"Task generation failed: {str(e)}")
