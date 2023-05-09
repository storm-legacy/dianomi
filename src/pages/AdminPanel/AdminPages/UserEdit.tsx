import React from 'react';
import { useParams } from 'react-router-dom';

function UserEdit() {
  console.log(useParams());
  const { UserId } = useParams();
  const userIdInt = UserId ? parseInt(UserId, 10) : undefined;
  return <p>{userIdInt}</p>;
}

export default UserEdit;
