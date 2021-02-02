# Cài đặt Postman https://www.postman.com/

# Lấy danh sách Sản Phẩm
curl -X GET \
  http://localhost:8080/product/list

# Thêm 1 sản phẩm vào danh sách
curl -X POST \
  http://localhost:8080/product/add \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -F id=1 \
  -F 'name=SP 1' \
  -F 'image=Ảnh 1'
  # Xóa 1 sản phẩm 
    http://localhost:8080/product/delete
