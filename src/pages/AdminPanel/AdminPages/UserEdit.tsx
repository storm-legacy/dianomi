import React from 'react';
import { useParams } from 'react-router-dom';

function UserEdit() {
  console.log(useParams());
  const { UserId } = useParams();
  const userIdInt = UserId ? parseInt(UserId, 10) : undefined;

  return <div> czy działą </div>;
}

export default UserEdit;
