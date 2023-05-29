import http from './axios.http';
import { Package } from './admin.service';

class ProfileService {
  GetPackage() {
    const controller = new AbortController();
    const request = http.get('/profile/package', { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }

  PostPayment() {
    const controller = new AbortController();
    const request = http.post('/profile/pay', {}, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
}
export default new ProfileService();
