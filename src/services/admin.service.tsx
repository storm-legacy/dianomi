import http from './axios.http';

interface VideoAddData {
  name: string;
  description: string;
  file_name: string;
  file_bucket: string;
  category_id: number | null;
  tags: string[];
}
interface UserEditData {
  email: string;
  verified: boolean;
  reset_password: boolean;
  packages: string[];
}

class AdminService {
  sendVideo(videoAdd: VideoAddData) {
    const controller = new AbortController();
    const request = http.post('/video', videoAdd, { signal: controller.signal });
    console.log(request);
    return { request, cancel: () => controller.abort() };
  }
  takeUser() {
    const controller = new AbortController();
    const request = http.get('/users', { signal: controller.signal });
    console.log(request);
    return { request, cancel: () => controller.abort() };
  }
  deleteUser(userId: number) {
    const controller = new AbortController();
    const request = http.delete('/users/' + userId, { signal: controller.signal });
    console.log(request);
    return { request, cancel: () => controller.abort() };
  }
  patchUser(userId: number | null, UserEdit: UserEditData) {
    const controller = new AbortController();
    const request = http.patch('/users/' + userId, UserEdit, { signal: controller.signal });
    console.log(request);
    return { request, cancel: () => controller.abort() };
  }
}

export default new AdminService();
export type { VideoAddData };
export type { UserEditData };
