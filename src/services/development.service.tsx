import http from './axios.http';
class DevelopmentService {
  GivePackageSelf(email: string) {
    const controller = new AbortController();
    const request = http.post('/dev/setpackage',
      {
        email: email,
        role: "premium",
        valid_for: 30
      }, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }

}
export default new DevelopmentService();
