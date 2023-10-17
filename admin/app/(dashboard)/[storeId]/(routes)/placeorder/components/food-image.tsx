"use client";

import Image from "next/image";
import { useEffect, useState } from "react";

const FoodsItemImage = () => {
  const [image, setImage] = useState("");
  const [loading, setLoading] = useState(false);

  // useEffect(() => {
  //   const fetchImage = async () => {
  //     setLoading(true);
  //     const res = await fetch("https://picsum.photos/180/180");
  //     setImage(res.url);
  //     setLoading(false);
  //   };
  //   fetchImage();
  // }, []);

  return (
    <div>
      <img
        src={image}
        loading={loading ? "lazy" : "eager"}
        alt="food image"
        className="w-[180px] h-[180px] rounded object-cover"
      />
    </div>
  );
};

export default FoodsItemImage;
