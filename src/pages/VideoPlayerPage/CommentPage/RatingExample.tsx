import React, { useEffect, useState } from 'react';
import { BsHandThumbsDownFill, BsHandThumbsDown, BsHandThumbsUpFill, BsHandThumbsUp } from 'react-icons/bs';
import videoService from '../../../services/video.service';
import { useParams } from 'react-router-dom';
const RatingExample = () => {
  const [rating, setRating] = useState<number>(0);
  const [upDisable, setUpDisable] = useState<boolean>(false);
  const [downDisable, setDownDisable] = useState<boolean>(false);
  const { VideoId } = useParams();
  const VideoIdInt = VideoId ? parseInt(VideoId, 10) : undefined;
  const handleRatingUp = () => {
    const { request } = videoService.sendUpVote(VideoIdInt);
    request
      .then(() => {
        setUpDisable(true);
        setDownDisable(false);
      })
      .catch((err) => {
        console.error(err);
      });
  };
  const handleRatingDown = () => {
    const { request } = videoService.sendDownVote(VideoIdInt);
    request
      .then(() => {
        setUpDisable(false);
        setDownDisable(true);
      })
      .catch((err) => {
        console.error(err);
      });
  };

  return (
    <div>
      <button onClick={handleRatingUp} className="btn btn-outline-primary" disabled={upDisable}>
        {!upDisable ? <BsHandThumbsUp /> : <BsHandThumbsUpFill />}
      </button>
      <button onClick={handleRatingDown} className="btn btn-outline-primary" disabled={downDisable}>
        {!downDisable ? <BsHandThumbsDown /> : <BsHandThumbsDownFill />}{' '}
      </button>
    </div>
  );
};

export default RatingExample;
