import React from 'react';
import GalleryCont from './GalleryCont';
import "./Gallery.css"

function GalleryRow() {
    
    return (
        <div class="gallery-row" >
            <GalleryCont/>
            <GalleryCont/>
            <GalleryCont/>
        </div>
    );
}

export default GalleryRow;