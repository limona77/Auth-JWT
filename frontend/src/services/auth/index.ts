import { AxiosResponse } from "axios";

import { AuthResponse } from "../../models/auth";
import { httpInstance } from "../../http";

export class AuthService {
  static async register(
    email: string,
    password: string,
  ): Promise<AxiosResponse<AuthResponse>> {
    return httpInstance.post<AuthResponse>("auth/register", {
      email,
      password,
    });
  }

  static async login(
    email: string,
    password: string,
  ): Promise<AxiosResponse<AuthResponse>> {
    return httpInstance.post<AuthResponse>("auth/login", {
      email,
      password,
    });
  }
}
