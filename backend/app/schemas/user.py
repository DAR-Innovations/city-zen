from pydantic import BaseModel

class UserCreate(BaseModel):
    email: str
    password: str
    name: str

class UserResponse(BaseModel):
    id: int
    email: str
    name: str
    is_active: bool

    class Config:
        orm_mode = True
