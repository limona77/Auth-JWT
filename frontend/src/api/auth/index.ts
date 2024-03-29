import { AuthService } from "../../services/auth";
import { IUser } from "../../models/User.ts";

export const fetchGetUser = async (): Promise<IUser> => {
  const response = await AuthService.getUser();
  return response.data;
};
