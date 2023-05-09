import React from 'react';
import { Link } from 'react-router-dom';
export const VideoList = () => {
  const VideoListItem = [
    {
      id: '1',
      name: 'wideo1',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi gravida massa mauris, id ',
      author: 'domino jachaś',
      tag: 'Lorem',
      categori: 'Lorem',
      file: 'Lorem/Lorem',
    },
    {
      id: '2',
      name: 'wideo2',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi gravida massa mauris, id  ',
      author: 'Lorens',
      tag: 'Lorem',
      categori: 'Lorem',
      file: 'Lorem/Lorem',
    },
    {
      id: '3',
      name: 'wideo3',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi gravida massa mauris, ',
      author: 'Lorens',
      tag: 'Lorem',
      categori: 'Lorem',
      file: 'Lorem/Lorem',
    },
    {
      id: '4',
      name: 'wideo4',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi gravida massa mauris,  ',
      author: 'Lorens',
      tag: 'Lorem',
      categori: 'Lorem',
      file: 'Lorem/Lorem',
    },
  ];
  return (
    <>
      <div className="text-center">
        <div className=" position-absolute top-50 start-50 translate-middle">
          <div className="row myRow myRowCont overflow-auto">
            <h1>Video List</h1>
            <div className="row myRow">
              <div className="col-1 border border-primary">ID</div>
              <div className="col-2 border border-primary">Name</div>
              <div className="col-2  border border-primary">Author</div>
              <div className="col-2 border border-primary">Description</div>
              <div className="col-1  border border-primary">Tag</div>
              <div className="col-2  border border-primary">Category</div>
              <div className="col-2 border border-primary"> Settings</div>
            </div>
            {VideoListItem.map((item, index) => (
              <div className="row" key={index}>
                <div className="col-1 border border-primary">{item.id}</div>
                <div className="col-2 border border-primary">{item.name}</div>
                <div className="col-2  border border-primary">{item.author}</div>
                <div className="col-2 border border-primary">{item.description}</div>
                <div className="col-1  border border-primary">{item.tag}</div>
                <div className="col-2  border border-primary">{item.categori}</div>
                <div className="col-2 border border-primary">
                  {' '}
                  Delete <Link to={'/VideoEdit/' + item.id}>Edit</Link> Ban
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </>
  );
};
