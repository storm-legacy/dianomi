import React, { useContext, useEffect, useState } from 'react';
import { BsHandThumbsDownFill, BsHandThumbsDown, BsHandThumbsUpFill, BsHandThumbsUp } from 'react-icons/bs';
import videoService from '../../../services/video.service';
import { useParams } from 'react-router-dom';
import profileService from '../../../services/profile.service';
import { AuthContext } from '../../../context/AuthContext';
const RatingExample = () => {
  const { user } = useContext(AuthContext);
  const [rating, setRating] = useState<number>(0);
  const [upDisable, setUpDisable] = useState<boolean>(false);
  const [downDisable, setDownDisable] = useState<boolean>(false);
  const { VideoId } = useParams();
  const VideoIdInt = VideoId ? parseInt(VideoId, 10) : undefined;
  useEffect(() => {
    const { request } = profileService.GetUserReaction({ email: user?.email, video_id: VideoIdInt });
    request.then((res) => {
      if (res.data.Value === 'down') {
        setUpDisable(false);
        setDownDisable(true);
      } else if (res.data.Value === 'up') {
        setUpDisable(true);
        setDownDisable(false);
      }
    });
  }, []);
  const handleRatingUp = () => {
    const { request } = videoService.sendUpVote(VideoIdInt);
    request
      .then(() => {
        const { request } = profileService.PostUserReaction({ email: user?.email, video_id: VideoIdInt, value: 'up' });
        request.catch((err) => console.error(err));
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
        const { request } = profileService.PostUserReaction({
          email: user?.email,
          video_id: VideoIdInt,
          value: 'down',
        });
        request.catch((err) => console.error(err));
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
