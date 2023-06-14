import { CommentReport } from '../types/comment.type';
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
interface VideoPatchData {
  name: string;
  description: string;
  category_id: number | null;
  is_premium: boolean | undefined;
  Tags: string[];
}
interface ThumbnailsData {
  video_id: number;
}
interface CategoriesName {
  name: string;
}
interface CommentData {
  email: string | undefined;
  video_id: number | undefined;
  comment: string;
}

class VideoService {
  sendVideo(videoAdd: VideoAddData) {
    const controller = new AbortController();
    const request = http.post('/video', videoAdd, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  editVideo(videoId: number | undefined, videoEdit: VideoPatchData) {
    const controller = new AbortController();
    const request = http.patch('/video/' + videoId, videoEdit, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  deleteVideo(videoId: number | undefined) {
    const controller = new AbortController();
    const request = http.delete('/video/' + videoId, { signal: controller.signal });
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
  takeSearchVideo(search: string) {
    const encoded = encodeURI(search);
    const controller = new AbortController();
    const request = http.get(`/video/search?phrase=${encoded}`, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  takeVideoRecommended(offset = 0, limit = 6) {
    const controller = new AbortController();
    const request = http.get(`/video?offset=${offset}&limit=${limit}`, { signal: controller.signal });
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
  sendComment(data: CommentData) {
    const controller = new AbortController();
    const request = http.post('/video/comment/', data, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  takeCommentVideoId(videoId: number | undefined) {
    const controller = new AbortController();
    const request = http.get('/video/comment/' + videoId, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  takeAllComment() {
    const controller = new AbortController();
    const request = http.get('/video/comment/all', { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  sendUpVote(videoId: number | undefined) {
    const controller = new AbortController();
    const request = http.post('/video/up/' + videoId, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  sendDownVote(videoId: number | undefined) {
    const controller = new AbortController();
    const request = http.post('/video/down/' + videoId, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  deleteComment(commentId: number | undefined) {
    const controller = new AbortController();
    const request = http.delete(`/video/comment/${commentId}`, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  reportComment(commentId: number, message: string) {
    const controller = new AbortController();
    const request = http.post(`/video/comment/report/${commentId}?message=${message}`,
      {}, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  getReportsForComment(commentId: number) {
    const controller = new AbortController();
    const request = http.get<CommentReport[]>(`/video/comment/report/${commentId}`,
      { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
  closeCommentReport(commentId: number) {
    const controller = new AbortController();
    const request = http.post(`/video/comment/close-report/${commentId}`, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }
}

export default new VideoService();
export type { VideoAddData };
export type { VideoPatchData };
export type { CategoriesName };
export type { CommentData };
