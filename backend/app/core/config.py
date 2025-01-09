from pydantic_settings import BaseSettings
from pydantic import Field
from typing import Optional

class Settings(BaseSettings):
    DATABASE_URL: str = Field(..., env="DATABASE_URL")
    USER_SECRET_KEY: str = Field(..., env="USER_SECRET_KEY")
    EMPLOYEE_SECRET_KEY: str = Field(..., env="EMPLOYEE_SECRET_KEY")

    ALGORITHM: str = Field(default="HS256", env="ALGORITHM")
    ACCESS_TOKEN_EXPIRE_MINUTES: int = Field(default=30, env="ACCESS_TOKEN_EXPIRE_MINUTES")

    UPLOAD_FOLDER: str = Field(default="uploads/", env="UPLOAD_FOLDER")
    ENVIRONMENT: str = Field(default="development", env="ENVIRONMENT")
    DEBUG: bool = Field(default=True, env="DEBUG")

    AI_MODEL_PATH: Optional[str] = Field(default=None, env="AI_MODEL_PATH")
    AI_CONFIG_PATH: Optional[str] = Field(default=None, env="AI_CONFIG_PATH")

    LOG_LEVEL: str = Field(default="info", env="LOG_LEVEL")

    DB_USER: str = Field(..., env="DB_USER")
    DB_PASSWORD: str = Field(..., env="DB_PASSWORD")
    DB_NAME: str = Field(..., env="DB_NAME")
    PG_EMAIL: str = Field(..., env="PG_EMAIL")
    PG_PASSWORD: str = Field(..., env="PG_PASSWORD")
    CLIENT_URL: str = Field(..., env="CLIENT_URL")
    SERVER_PORT: int = Field(default=8000, env="SERVER_PORT")

    class Config:
        env_file = ".env"
        env_file_encoding = "utf-8"

settings = Settings()
