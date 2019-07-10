

$(document).ready(()=> {
    let raf= window.requestAnimationFrame?window.requestAnimationFrame:function(callback){return setTimeout(callback,12)};
    let $bulin = $('div.sprite-bulin');
        if ($bulin[0]) {
            $bulin.redraw = function () {$bulin.css('transform', 'translate3d(' + Math.floor($bulin.X) + 'px, ' + Math.floor($bulin.Y) + 'px, 0) ' + (!$bulin.reversed ? 'rotateY(180deg)' : ''));};
            let bulin_width = 130;
            let range = {minX: -50, maxX: document.body.clientWidth - bulin_width + 50, minY: -50, maxY: window.innerHeight - bulin_width + 50};
            $bulin.speedX = (Math.random() + 1) * 2;
            $bulin.speedY = (Math.random() + 1) * 2;
            $bulin.X = -130;
            $bulin.Y = 0;
            function bulinMove() {
                if ($bulin.X < range.minX) {
                    $bulin.speedX = Math.abs($bulin.speedX);
                    $bulin.reversed = !1;
                } else if ($bulin.X > range.maxX) {
                    $bulin.speedX = -$bulin.speedX;
                    $bulin.reversed = !0;
                }
                $bulin.speedY = $bulin.Y < range.minY ? Math.abs($bulin.speedY) : $bulin.Y > range.maxY ? -$bulin.speedY : $bulin.speedY;
                $bulin.X += $bulin.speedX;
                $bulin.Y += $bulin.speedY;
                $bulin.redraw();
                return raf(bulinMove)
            }
            setTimeout(() => {
                bulinMove();
            }, 1000);
        }
    }
);

