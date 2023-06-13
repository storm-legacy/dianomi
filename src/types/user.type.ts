export interface User {
  email: string;
  role: string;
  verified: boolean;
  authToken?: string;
  banned: boolean;
}
