import React, { useEffect, useState } from 'react';
import ReactPlayer from 'react-player';
import videoService from '../../services/video.service';
import { useParams } from 'react-router-dom';
import { Notify } from 'notiflix';

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

export const VideoPlayer = () => {
  const [dataVideo, setDataVideo] = useState<VideoData[]>([]);
  const [videoTags, setVideoTags] = useState<string[]>([]);
  const [videoName, setVideoName] = useState('');
  const [videoThumbnail, setVideoThumbnail] = useState('');
  const [selectedOption, setSelectedOption] = useState(0);
  const [recommendedVideos, setRecommendedVideos] = useState<VideoItemData[]>([]);

  const { VideoId } = useParams();
  const VideoIdInt = VideoId ? parseInt(VideoId, 10) : undefined;

  useEffect(() => {
    const { request, cancel } = videoService.takeVideoRecommended();
    request
      .then(({ data }) => {
        setRecommendedVideos(
          data.filter((vid: VideoItemData) => {
            return vid.id != VideoIdInt;
          }),
        );
      })
      .catch((err) => {
        console.log(err.message);
      });

    return () => cancel();
  }, []);

  useEffect(() => {
    const { request } = videoService.takeVideoId(VideoIdInt);
    request
      .then((res) => {
        setVideoName(res.data.name);
        setDataVideo(res.data.videos);
        setVideoThumbnail(res.data.thumbnail_url);
        setVideoTags(res.data.tags);
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
    <>
      <div className="container-fluid mx-0 py-8">
        <div className="container py-4">
          <h5>Polecane materiały:</h5>
          <div className="row" style={{ maxHeight: '90vh' }}>
            <div className="col-4 overflow-y-visible">
              {recommendedVideos.map((item, index) => {
                return (
                  <a href={`/VideoPlayer/${item.id}`} key={index}>
                    <div className="card text-bg-secondary mb-3" key={index} style={{ maxWidth: '18rem' }}>
                      <div className="card-header">{item.name}</div>
                      <div className="card-body">
                        <p className="card-text">{item.description.substring(0, 64)}...</p>
                      </div>
                    </div>
                  </a>
                );
              })}
            </div>
            <div className="col-8 d-flex flex-column align-items-start">
              <div className="col-12 p-2 d-flex justify-content-start">
                <div className="col-2">
                  <select className="form-select" value={selectedOption} onChange={handleSelectChange}>
                    <option value={0}>720p</option>
                    <option value={2}>480p</option>
                    <option value={1}>360p</option>
                  </select>
                </div>
              </div>
              <div className="col-12">
                <ReactPlayer
                  light={videoThumbnail}
                  className=""
                  url={dataVideo[selectedOption]?.file_path}
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
              <div className="col-12">
                <h3 className="p-2">Lorem ipsum dolorem mi</h3>
              </div>
              <div className="col-12 px-2 d-flex flex-row justify-content-start">
                <strong>Tags: </strong>
                {videoTags.map((item, index) => {
                  return (
                    <p className="px-2" key={index}>
                      <i>test</i>
                    </p>
                  );
                })}
              </div>
              <div className="p-2 col-8">
                <span>
                  Lorem ipsum dolor sit, amet consectetur adipisicing elit. Minus ad inventore dicta minima! Enim
                  laboriosam ipsum sit explicabo ad labore eos, at facere quia ipsa fugiat qui distinctio eaque autem!
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};
