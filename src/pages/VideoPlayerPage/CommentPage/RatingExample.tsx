import React, { useEffect, useState } from 'react';
import { BsHandThumbsDownFill, BsHandThumbsDown, BsHandThumbsUpFill, BsHandThumbsUp } from 'react-icons/bs';
const RatingExample = () => {
  const [rating, setRating] = useState<number>(0);
  const [upDisable, setUpDisable] = useState<boolean>(false);
  const [downDisable, setDownDisable] = useState<boolean>(false);

  const handleRatingUp = () => {
    console.log('przed', rating);
    setRating(+1);
    setUpDisable(true);
    setDownDisable(false);
    console.log(rating);
  };
  const handleRatingDown = () => {
    console.log('przed', rating);
    setRating(-1);
    setUpDisable(false);
    setDownDisable(true);
    console.log(rating);
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
