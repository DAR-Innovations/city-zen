from fastapi import FastAPI, HTTPException
from ai import generate_task_metadata, validate_task_completion
from db import fetch_departments_and_task_types, get_db_connection
from models import GenerateTaskRequest, ValidateTaskRequest, TaskMetadata, ValidateTaskResponse

app = FastAPI()

try:
    db_connection = get_db_connection()
    db_connection.close()
    print("Database connection established successfully.")
except Exception as e:
    print(f"Failed to connect to the database: {e}")
    raise RuntimeError(f"Database connection failed: {e}")


@app.post("/generate-task-metadata/", response_model=TaskMetadata)
def generate_task_metadata_endpoint(request: GenerateTaskRequest):
    departments = fetch_departments_and_task_types()
    try:
        result = generate_task_metadata(request.image_url, request.issue_description, departments)
        return result
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@app.post("/validate-task-completion/", response_model=ValidateTaskResponse)
def validate_task_completion_endpoint(request: ValidateTaskRequest):
    try:
        result = validate_task_completion(request.original_image_url, request.completed_image_url)
        return result
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
