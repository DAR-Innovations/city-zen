from dns.dnssec import RSASHA256
from fastapi import Request, HTTPException
from jose import JWTError, jwt
from functools import wraps
from app.core.config import settings


def verify_token(token: str) -> dict:
    try:
        payload = jwt.decode(token, settings.EMPLOYEE_SECRET_KEY, algorithms=settings.ALGORITHM)
        return payload
    except JWTError:
        raise HTTPException(status_code=401, detail="Invalid employee token")


def require_auth(func):
    @wraps(func)
    async def wrapper(*args, request: Request, **kwargs):
        token = request.cookies.get("EMPLOYEE_ACCESS_TOKEN")
        if not token:
            raise HTTPException(status_code=401, detail="No employee token provided")

        payload = verify_token(token)
        request.state.employee= payload
        return await func(*args, request=request, **kwargs)

    return wrapper