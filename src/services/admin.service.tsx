import http from './axios.http';

interface VideoAddData {
  name: string;
  description: string;
  file_name: string;
  file_bucket: string;
  category_id: number | null;
  tags: string[];
}

class AdminService {
  sendVideo(videoAdd: VideoAddData) {
    const controller = new AbortController();
    const request = http.post('/admin/video', videoAdd, { signal: controller.signal });
    console.log(request);
    return { request, cancel: () => controller.abort() };
  }
}

export default new AdminService();
export type { VideoAddData };
