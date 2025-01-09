from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.future import select
from fastapi import HTTPException
from passlib.context import CryptContext
from datetime import datetime, timedelta
from jose import JWTError, jwt
from app.models.agents.users import Users
from app.models.agents.employees import Employees
from app.core.config import settings
from app.modules.auth.types.auth_dtos import LoginRequest, RegisterRequest, LoginResponse

pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")


def verify_password(plain_password: str, hashed_password: str) -> bool:
    return pwd_context.verify(plain_password, hashed_password)


def create_access_token(data: dict, expires_in: timedelta, secret_key: str = settings.USER_SECRET_KEY) -> str:
    expire = datetime.now() + expires_in
    to_encode = data.copy()
    to_encode.update({"exp": expire})
    encoded_jwt = jwt.encode(to_encode, secret_key, algorithm=settings.ALGORITHM)
    return encoded_jwt


def get_password_hash(password: str) -> str:
    return pwd_context.hash(password)


def decode_token(token: str, secret_key: str = settings.USER_SECRET_KEY) -> dict:
    try:
        payload = jwt.decode(token, secret_key, algorithms=[settings.ALGORITHM])
        return payload
    except JWTError:
        raise HTTPException(status_code=401, detail="Invalid token or expired")


def create_employee_token(employee: Employees) -> str:
    employee_claims = {
        "id": employee.id,
        "department_id": employee.department_id,
        "phone": employee.phone,
        "isVerified": employee.is_verified,
        "role": employee.role
    }
    access_token = create_access_token(
        data=employee_claims,
        expires_in=timedelta(minutes=settings.ACCESS_TOKEN_EXPIRE_MINUTES),
        secret_key=settings.EMPLOYEE_SECRET_KEY
    )
    return access_token


def create_user_token(user: Users) -> str:
    user_claims = {
        "id": user.id,
        "phone": user.phone,
        "isVerified": user.is_verified,
        "role": user.role
    }
    access_token = create_access_token(
        data=user_claims,
        expires_in=timedelta(minutes=settings.ACCESS_TOKEN_EXPIRE_MINUTES)
    )
    return access_token


class AuthService:
    def __init__(self, db: AsyncSession):
        self.db = db

    async def user_signin(self, dto: LoginRequest) -> LoginResponse:

        async with self.db as session:
            result = await session.execute(select(Users).filter(Users.phone == dto.phone))
            user = result.scalars().first()

            if not user or not verify_password(dto.password, user.hashed_password):
                raise HTTPException(status_code=401, detail="Invalid user phone or password")

            token = create_user_token(user)
            return LoginResponse(access_token=token)

    async def user_signup(self, dto: RegisterRequest) -> int:
        async with self.db as session:
            result = await session.execute(select(Users).filter(Users.phone == dto.phone))
            existing_user = result.scalars().first()
            if existing_user:
                raise HTTPException(status_code=400, detail="User phone already registered")

            user = Users(
                name=dto.name,
                phone=dto.phone,
                hashed_password=get_password_hash(dto.password)
            )
            session.add(user)
            await session.commit()
            await session.refresh(user)
            return user.id

    async def employee_signin(self, dto: LoginRequest) -> LoginResponse:
        async with self.db as session:
            result = await session.execute(select(Employees).filter(Employees.phone == dto.phone))
            employee = result.scalars().first()
            if not employee or not verify_password(dto.password, employee.hashed_password):
                raise HTTPException(status_code=401, detail="Invalid employee phone or password")

            token = create_employee_token(employee)
            return LoginResponse(access_token=token)
