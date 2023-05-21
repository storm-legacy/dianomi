import React from 'react';
import { useParams } from 'react-router-dom';
import videoService from '../../../services/video.service';

export const VideoDelete = () => {
  const { VideoId } = useParams();
  const VideoIdInt = VideoId ? parseInt(VideoId, 10) : undefined;
  const { request } = videoService.deleteVideo(VideoIdInt);
  request
    .then((res) => {
      console.log(res);
    })
    .catch((err) => {
      console.error(err.message);
    });

  return <div>{VideoIdInt}</div>;
};
