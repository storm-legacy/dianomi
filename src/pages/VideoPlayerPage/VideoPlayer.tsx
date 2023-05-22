import React, { useEffect, useState } from 'react';
import ReactPlayer from 'react-player';
import videoService from '../../services/video.service';
import { useParams } from 'react-router-dom';
export const VideoPlayer = () => {
  interface VideoData {
    duration: number;
    file_path: string;
    resolution: string;
  }

  interface VideoItemData {
    id: number;
    name: string;
    description: string;
    category: string;
    tags: string[];
    thumbnail_url: string;
    videos: VideoData[];
  }
  const [dataVideo, setDataVideo] = useState<VideoData[]>([]);
  const [videoName, setVideoName] = useState('');
  const [videoThumbnail, setVideoThumbnail] = useState('');
  const [selectedOption, setSelectedOption] = useState(0);
  const { VideoId } = useParams();
  const VideoIdInt = VideoId ? parseInt(VideoId, 10) : undefined;
  console.log(VideoIdInt);
  useEffect(() => {
    const { request } = videoService.takeVideoId(VideoIdInt);
    request
      .then((res) => {
        console.log(res.data.videos);
        setVideoName(res.data.name);
        setDataVideo(res.data.videos);
        setVideoThumbnail(res.data.thumbnail_url);
        console.log(videoThumbnail);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  useEffect(() => {
    console.log('Wybrana opcja:', selectedOption);
  }, [selectedOption]);

  const handleSelectChange = (event: { target: { value: string } }) => {
    const value = parseInt(event.target.value, 10);
    setSelectedOption(value);
  };
  return (
    <div className="position-absolute top-50 start-50 translate-middle">
      <h3>{videoName}</h3>
      {dataVideo.length > 0 && (
        <div className="">
          <ReactPlayer
            light={videoThumbnail}
            className=""
            url={dataVideo[selectedOption].file_path}
            controls
            config={{
              file: {
                attributes: {
                  controlsList: 'nodownload',
                },
              },
            }}
          />
        </div>
      )}
      <select className="form-select" value={selectedOption} onChange={handleSelectChange}>
        <option value={0}>720p</option>
        <option value={2}>480p</option>
        <option value={1}>360p</option>
      </select>
    </div>
  );
};
