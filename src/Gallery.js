import React from 'react';
import GalleryRow from './GalleryRow';
import "./Gallery.css"

function Gallery() {
    
    return (
        <div class="gallery" >
            <GalleryRow/>
            <GalleryRow/>
            <GalleryRow/>
        </div>
    );
}

export default Gallery;