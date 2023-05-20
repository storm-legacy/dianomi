import http from './axios.http';

interface VideoAddData {
  name: string;
  description: string;
  file_name: string;
  file_bucket: string;
  thumbnail_name: string;
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
    return { request, cancel: () => controller.abort() };
  }
  sendCategori(categorieAdd: CategoriesName) {
    const controller = new AbortController();
    const request = http.post('/video/category', categorieAdd, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  takeVideo() {
    const controller = new AbortController();
    const request = http.get('/video/all', { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  takeVideoRecommended() {
    const controller = new AbortController();
    const request = http.get('/video', { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  takeVideoId(videoId: number | undefined) {
    const controller = new AbortController();
    const request = http.get('/video/' + videoId, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  takeCategori() {
    const controller = new AbortController();
    const request = http.get('/video/category/all', { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  sendThumbnails(sendThumbnails: ThumbnailsData) {
    const controller = new AbortController();
    const request = http.post('/admin/video', sendThumbnails, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
}

export default new VideoService();
export type { VideoAddData };
export type { CategoriesName };
