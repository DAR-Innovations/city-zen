from pydantic import BaseModel

class RegionTasksRequestDTO(BaseModel) :
    longitude: float
    latitude: float
    radius: float

class RegionTasksResponseDTO(BaseModel) :
    createdIssues: int
    resolvedTasks: int

class CompletedTasksDTO(BaseModel) :
    completedTasks: int




