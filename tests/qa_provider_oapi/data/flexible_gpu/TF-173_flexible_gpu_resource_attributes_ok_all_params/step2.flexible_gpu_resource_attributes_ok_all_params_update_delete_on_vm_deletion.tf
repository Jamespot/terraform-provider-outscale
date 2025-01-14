resource "outscale_flexible_gpu" "fGPU-1" {
   model_name             =  "nvidia-k2"
   generation             =  var.fgpu_gen
   subregion_name         =  "${var.region}a"
   delete_on_vm_deletion  =   false
}

