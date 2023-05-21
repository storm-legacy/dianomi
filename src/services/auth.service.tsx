import http from './axios.http';

class AuthService {
  connectionCheck() {
    const controller = new AbortController();
    const request = http.get('/', { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }

  login(email: string, password: string) {
    const controller = new AbortController();
    const request = http.post(
      '/auth/login',
      {
        email: email,
        password: password,
      },
      { signal: controller.signal },
    );
    return { request, cancel: () => controller.abort() };
  }

  register(email: string, password: string, passwordRepeat: string) {
    const controller = new AbortController();
    const request = http.post(
      '/auth/register',
      {
        email: email,
        password: password,
        password_repeat: passwordRepeat,
      },
      { signal: controller.signal },
    );
    return { request, cancel: () => controller.abort() };
  }

  refresh() {
    const controller = new AbortController();
    const request = http.post('/auth/refresh', {}, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }

  logout() {
    const controller = new AbortController();
    const request = http.post('/auth/logout', {}, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  genreset(email: string) {
    const controller = new AbortController();
    const request = http.post('/auth/genreset', { email: email }, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  resetGet(code: string) {
    const controller = new AbortController();
    const request = http.get('/auth/reset?validate=' + code, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  resetPost(validate: string, password: string, passwordRepeat: string) {
    const controller = new AbortController();
    const request = http.post(
      '/auth/reset',
      {
        validate: validate,
        password: password,
        password_repeat: passwordRepeat,
      },
      { signal: controller.signal },
    );
    return { request, cancel: () => controller.abort() };
  }
}

export default new AuthService();
