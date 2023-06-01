import http from './axios.http';
interface userResData {
  email: string | undefined;
  OldPassword: string;
}
class TestService {
  GetOldPassword(data: userResData) {
    const controller = new AbortController();
    const request = http.post('/test/', data, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
}
export default new TestService();
export type { userResData };
