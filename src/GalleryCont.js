
import React from 'react';
import "./Gallery.css"

class GalleryCont extends React.Component {

  constructor(props) {
    super(props);
    this.state = {isSelected: false, imageStyle: "gallery-image"};


    // This binding is necessary to make `this` work in the callback
    this.handleClick = this.handleClick.bind(this);
  }
  handleClick() {
    console.log("touched")
    this.setState(state => ({
      isSelected: !state.isSelected,
      imageStyle: this.state.isSelected ? "gallery-image-selected" : "gallery-image",
    }));
  }
    
  render() {
    return (
      <div class="gallery-cont">
          <img onClick={this.handleClick} class={this.state.imageStyle} src="https://images.unsplash.com/photo-1593642632823-8f785ba67e45?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1189&q=80"  alt="logo" />
      </div>
    );
  }
}

export default GalleryCont;


