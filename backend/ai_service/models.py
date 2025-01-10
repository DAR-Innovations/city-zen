from pydantic import BaseModel
from typing import Optional

class TaskMetadata(BaseModel):
    title: str
    description: str
    urgency: str
    complexity: str
    department: Optional[str] = None


class GenerateTaskRequest(BaseModel):
    image_url: str
    issue_description: str


class ValidateTaskRequest(BaseModel):
    original_image_url: str
    completed_image_url: str


class ValidateTaskResponse(BaseModel):
    is_valid: bool
    reason: str
