import http from './axios.http';
import { Package } from './admin.service';
interface emailData {
  ErrorTitle: string;
  ErrorDescription: string;
  ReportedBy: string;
}
interface oldPass {
  email: string | undefined;
  OldPassword: string;
}
interface userResData {
  email: string | undefined;
  NewPassword: string;
}
interface Email {
  email: string | undefined;
}
interface metricData {
  email: string | undefined;
  video_id: number | undefined;
  time_spent_watching: number | undefined;
  stopped_at: number | undefined;
}

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

  PostRaport(data: emailData) {
    const controller = new AbortController();
    const request = http.post('/report/', data, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  PostOldPassword(data: oldPass) {
    const controller = new AbortController();
    const request = http.post('/profile/comparison/', data, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  PostNewPassword(data: userResData) {
    const controller = new AbortController();
    const request = http.post('/profile/new/', data, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  GetUserVideoMetric(data: Email) {
    const controller = new AbortController();
    const request = http.post('/metrics/user/', data, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  PostVideoMertics(data: metricData) {
    const controller = new AbortController();
    const request = http.post('/metrics/', data, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
}
export default new ProfileService();
export type { emailData };
export type { oldPass };
export type { userResData };
export type { Email };
export type { metricData };
