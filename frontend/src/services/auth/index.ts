import { AxiosResponse } from "axios";

import { IAuthResponse } from "../../models/auth.ts";
import { httpInstance } from "../../http";
import { IUser } from "../../models/User.ts";

export class AuthService {
  static async register(
    email: string,
    password: string,
  ): Promise<AxiosResponse<IAuthResponse>> {
    return httpInstance.post<IAuthResponse>("auth/register", {
      email,
      password,
    });
  }

  static async login(
    email: string,
    password: string,
  ): Promise<AxiosResponse<IAuthResponse>> {
    return httpInstance.post<IAuthResponse>("auth/login", {
      email,
      password,
    });
  }

  static async me(): Promise<AxiosResponse<IUser>> {
    return httpInstance.get<IUser>("auth/me", {
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
    });
  }
}
