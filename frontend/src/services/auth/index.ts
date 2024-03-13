import { AxiosResponse } from "axios";

import { AuthResponse } from "../../models/auth";
import { axiosInstance } from "../../axios";

export class AuthService {
  static async register(
    email: string,
    password: string,
  ): Promise<AxiosResponse<AuthResponse>> {
    return axiosInstance.post<AuthResponse>("auth/register", {
      email,
      password,
    });
  }
}
