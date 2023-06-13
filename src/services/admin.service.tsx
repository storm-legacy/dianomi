import http from './axios.http';

interface VideoAddData {
  name: string;
  description: string;
  file_name: string;
  file_bucket: string;
  category_id: number | null;
  tags: string[];
}

export enum Tier {
  free = 'free',
  premium = 'premium',
  administrator = 'administrator',
}

interface Package {
  id: number;
  user_id: number;
  tier: Tier;
  created_at: Date;
  valid_from: string;
  valid_until: string;
}
interface UserEditData {
  email: string;
  verified: boolean;
  reset_password: boolean;
}

class AdminService {
  sendVideo(videoAdd: VideoAddData) {
    const controller = new AbortController();
    const request = http.post('/video', videoAdd, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }

  takeUser() {
    const controller = new AbortController();
    const request = http.get('/users', { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }

  getUser(userID: number) {
    const controller = new AbortController();
    const request = http.get(`/users/${userID}`, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }

  deleteUser(userId: number) {
    const controller = new AbortController();
    const request = http.delete('/users/' + userId, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }

  patchUser(userId: number | null, UserEdit: UserEditData) {
    const controller = new AbortController();
    const request = http.patch('/users/' + userId, UserEdit, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }

  getUserPackages(userId: number, limit = 5, page = 1) {
    const controller = new AbortController();
    const request = http.get<Package[]>(`/users/packages/${userId}?page=${page}&limit=${limit}`, {
      signal: controller.signal,
    });
    return { request, cancel: () => controller.abort() };
  }

  deleteUserPackage(packageId: number) {
    const controller = new AbortController();
    const request = http.delete(`/users/packages/pack/${packageId}`, { signal: controller.signal });
    return { request, cancel: () => controller.abort() };
  }

  postUserPackage({ user_id, tier, valid_from, valid_until }: Package) {
    const controller = new AbortController();
    const request = http.post(
      `/users/packages`,
      {
        user_id: user_id,
        tier: tier,
        valid_from: valid_from,
        valid_until: valid_until,
      },
      { signal: controller.signal },
    );
    return { request, cancel: () => controller.abort() };
  }

  patchUserPackage({ id, user_id, tier, valid_from, valid_until }: Package) {
    const controller = new AbortController();
    const request = http.patch(
      `/users/packages/pack/id`,
      {
        id: id,
        user_id: user_id,
        tier: tier,
        valid_from: valid_from,
        valid_until: valid_until,
      },
      { signal: controller.signal },
    );
    return { request, cancel: () => controller.abort() };
  }

  postBanUser(id: number) {
    const controller = new AbortController();
    const request = http.post(
      `/users/ban/${id}`,
      { signal: controller.signal },
    );
    return { request, cancel: () => controller.abort() };
  }

  postUnbanUser(id: number) {
    const controller = new AbortController();
    const request = http.post(
      `/users/unban/${id}`,
      { signal: controller.signal },
    );
    return { request, cancel: () => controller.abort() };
  }
}

export default new AdminService();
export type { VideoAddData, UserEditData, Package };
