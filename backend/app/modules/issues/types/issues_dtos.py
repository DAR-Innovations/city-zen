from pydantic import BaseModel

class PostIssueRequestDTO(BaseModel):
    name: str
    description: str
    longitude: float
    latitude: float

class IssueResponseDTO(BaseModel):
    id: int
    name: str
    description: str
    longitude: float
    latitude: float
    is_completed: bool
    created_at: str
    image_url: str
    author_id: int