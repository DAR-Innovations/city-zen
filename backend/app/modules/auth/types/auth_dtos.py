from pydantic import BaseModel

class RegisterRequest(BaseModel):
    name: str
    phone: str
    password: str

class LoginRequest(BaseModel):
    phone: str
    password: str

class LoginResponse(BaseModel):
    access_token: str