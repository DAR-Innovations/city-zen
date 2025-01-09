from fastapi import APIRouter, Depends, Response, HTTPException

from app.modules.auth.auth_service import AuthService
from app.modules.auth.types.auth_dtos import LoginResponse, LoginRequest, RegisterRequest

router = APIRouter()

@router.post("/user/signin")
async def user_signin(data: LoginRequest, service: AuthService = Depends(), response: Response = Depends()
    ) -> LoginResponse:
    try:
        login_response = await service.user_signin(data)

        response.set_cookie(key="ACCESS_TOKEN", value=login_response.access_token, httponly=True, secure=True)

        return login_response

    except HTTPException as e:
        raise e
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))

@router.post("/user/signup")
async def user_signup(data: RegisterRequest, service: AuthService = Depends()):
    try:
        user_id = await service.user_signup(data)
        return {"message": "Registration successful", "user_id": user_id}
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))

@router.post("/employee/signin", response_model=None)
async def employee(data: LoginRequest, service: AuthService = Depends(), response: Response = Depends()
    ) -> LoginResponse:

    try:
        login_response = await service.employee_signin(data)

        response.set_cookie(key="ACCESS_TOKEN", value=login_response.access_token, httponly=True, secure=True)

        return login_response
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))