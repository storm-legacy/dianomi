import http from './axios.http';

interface VideoAddData {
  name: string;
  description: string;
  file_name: string;
  file_bucket: string;
  category_id: number | null;
  tags: string[];
}
interface ThumbnailsData {
  video_id: number;
}
interface CategoriesName {
  name: string;
}

class VideoService {
  sendVideo(videoAdd: VideoAddData) {
    const controller = new AbortController();
    const request = http.post('/video', videoAdd, { signal: controller.signal });
    console.log(request);
    return { request, cancel: () => controller.abort() };
  }
  sendCategori(categorieAdd: CategoriesName) {
    const controller = new AbortController();
    const request = http.post('/video/category', categorieAdd, { signal: controller.signal });
    console.log(request);
    return { request, cancel: () => controller.abort() };
  }
  takeVideo() {
    const controller = new AbortController();
    const request = http.get('/video/all', { signal: controller.signal });
    console.log(request);
    return { request, cancel: () => controller.abort() };
  }
  takeCategori() {
    const controller = new AbortController();
    const request = http.get('/video/category', { signal: controller.signal });
    console.log(request);
    return { request, cancel: () => controller.abort() };
  }
  sendThumbnails(sendThumbnails: ThumbnailsData) {
    const controller = new AbortController();
    const request = http.post('/admin/video', sendThumbnails, { signal: controller.signal });
    console.log(request);
    return { request, cancel: () => controller.abort() };
  }
}

export default new VideoService();
export type { VideoAddData };
export type { CategoriesName };
