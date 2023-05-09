import React, { useState } from 'react';
import { useParams } from 'react-router-dom';
export const VideoEdit = () => {
  const { VideoId } = useParams();
  const VideoIdInt = VideoId ? parseInt(VideoId, 10) : undefined;
  console.log(VideoIdInt);
  const [videoName, setVideoName] = useState('');
  const [videoFile, setVideoFile] = useState('');
  const [videoAuthor, setVideoAuthor] = useState('');
  const [videoCategori, setVideoCategori] = useState('');
  const [videoDiscription, setVideoDiscription] = useState('');
  const [videoTag, setVideoTag] = useState('');
  const VideoListItem = [
    {
      id: 1,
      name: 'wideo1',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi gravida massa mauris, id ',
      author: 'domino jachaś',
      tag: 'Lorem',
      categori: 'Lorem',
      file: 'Lorem/Lorem',
    },
    {
      id: 2,
      name: 'wideo2',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi gravida massa mauris, id  ',
      author: 'Lorens',
      tag: 'Lorem',
      categori: 'Lorem',
      file: 'Lorem/Lorem',
    },
    {
      id: 3,
      name: 'wideo3',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi gravida massa mauris, ',
      author: 'Lorens',
      tag: 'Lorem',
      categori: 'Lorem',
      file: 'Lorem/Lorem',
    },
    {
      id: 4,
      name: 'wideo4',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi gravida massa mauris,  ',
      author: 'Lorens',
      tag: 'Lorem',
      categori: 'Lorem',
      file: 'Lorem/Lorem',
    },
  ];
  const handleSubmit = (event: any) => {
    event.preventDefault();

    const videoData = {
      videoName: videoName,
      videoDiscription: videoDiscription,
      videoAuthor: videoAuthor,
      videoTag: videoTag,
      videoCategori: videoCategori,
      videoFile: videoFile,
    };
  };
  return (
    <div>
      {VideoListItem.map((item) => {
        if (item.id === VideoIdInt) {
          return (
            <div
              key={item.id}
              className="position-absolute top-50 start-50 translate-middle text-center float-start shadow-lg p-3 mb-5 bg-white rounded"
            >
              <h3>Edit Video</h3>
              <p></p>
              <form onSubmit={handleSubmit} className="row">
                <label>
                  <p>Name</p>
                  <input
                    className="form-control"
                    type="text"
                    value={item.name}
                    onChange={(event) => setVideoName(event.target.value)}
                  />
                </label>
                <label>
                  <p>Discription</p>
                  <input
                    className="form-control"
                    type="text"
                    value={item.description}
                    onChange={(event) => setVideoDiscription(event.target.value)}
                  />
                </label>
                <label>
                  <p>File</p>
                  <input
                    className="form-control"
                    type="file"
                    value={videoFile}
                    onChange={(event) => setVideoFile(event.target.value)}
                  />
                </label>
                <label>
                  <p>Author</p>
                  <input
                    className="form-control"
                    type="text"
                    value={item.author}
                    onChange={(event) => setVideoAuthor(event.target.value)}
                  />
                </label>
                <label>
                  <p>Category</p>
                  <input
                    className="form-control"
                    type="list"
                    list="CategoryList"
                    value={item.categori}
                    onChange={(event) => setVideoCategori(event.target.value)}
                  />
                  <datalist id="CategoryList">
                    <option value="C++" />
                    <option value="Go" />
                    <option value="HowTo" />
                  </datalist>
                </label>
                <label>
                  <p>Tag</p>
                  <input
                    className="form-control"
                    list="TagList"
                    value={item.tag}
                    onChange={(event) => setVideoTag(event.target.value)}
                  />
                  <datalist id="TagList">
                    <option value="1" />
                    <option value="2" />
                    <option value="3" />
                  </datalist>
                </label>
                <p></p>
                <button type="submit" className="btn btn-primary ">
                  Edit
                </button>
              </form>
            </div>
          );
        } else {
          return null;
        }
      })}
    </div>
  );
};
