import { IUser } from "./User.ts";

export interface IAuthResponse {
  accessToken: string;
  refreshToken: string;
  UserData: IUser;
}
